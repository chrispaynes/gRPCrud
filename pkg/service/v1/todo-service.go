package v1

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/chrispaynes/gRPCrud/pkg/api/v1"
)

const apiVersion = "v1"

// TodoServiceServer ...
type TodoServiceServer struct {
	db *sqlx.DB
}

// NewTodoServiceServer creates a Todo service
func NewTodoServiceServer(db *sqlx.DB) v1.TodoServiceServer {
	return &TodoServiceServer{db: db}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *TodoServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// connect returns SQL database connection from the pool
func (s *TodoServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database "+err.Error())
	}
	return c, nil
}

// Create new Todo task
func (s *TodoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	// if err := s.checkAPI(req.Api); err != nil {
	// 	return nil, err
	// }

	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	reminder, err := ptypes.Timestamp(req.Todo.Reminder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	}

	// insert Todo entity data
	var id int64
	c.QueryRowContext(ctx, `INSERT INTO todo (title, description, reminder) VALUES ($1, $2, $3) RETURNING todo_id`,
		req.Todo.Title, req.Todo.Description, reminder).Scan(&id)

	if id == 0 {
		return nil, status.Error(codes.Unknown, "failed to insert into todo")
	}

	return &v1.CreateResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// ReadAll reads all
func (s *TodoServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	// if err := s.checkAPI(req.Api); err != nil {
	// 	return nil, err
	// }

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// get Todo list
	rows, err := c.QueryContext(ctx, "select todo_id, title, description, reminder from todo")

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from Todo "+err.Error())
	}
	defer rows.Close()

	var reminder time.Time
	list := []*v1.Todo{}
	for rows.Next() {
		td := new(v1.Todo)
		if err := rows.Scan(&td.TodoId, &td.Title, &td.Description, &reminder); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from Todo row-> "+err.Error())
		}
		td.Reminder, err = ptypes.TimestampProto(reminder)
		if err != nil {
			return nil, status.Error(codes.Unknown, "reminder field has invalid format-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from Todo-> "+err.Error())
	}

	return &v1.ReadAllResponse{
		Api:   apiVersion,
		Todos: list,
	}, nil
}
