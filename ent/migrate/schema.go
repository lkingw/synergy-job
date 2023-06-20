// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SysTasksColumns holds the columns for the "sys_tasks" table.
	SysTasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "name", Type: field.TypeString},
		{Name: "task_group", Type: field.TypeString},
		{Name: "cron_expression", Type: field.TypeString},
		{Name: "pattern", Type: field.TypeString},
		{Name: "payload", Type: field.TypeString},
	}
	// SysTasksTable holds the schema information for the "sys_tasks" table.
	SysTasksTable = &schema.Table{
		Name:       "sys_tasks",
		Columns:    SysTasksColumns,
		PrimaryKey: []*schema.Column{SysTasksColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "task_pattern",
				Unique:  true,
				Columns: []*schema.Column{SysTasksColumns[7]},
			},
		},
	}
	// SysTaskLogsColumns holds the columns for the "sys_task_logs" table.
	SysTaskLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "started_at", Type: field.TypeTime, Comment: "Task Started Time | 任务启动时间"},
		{Name: "finished_at", Type: field.TypeTime, Comment: "Task Finished Time | 任务完成时间"},
		{Name: "result", Type: field.TypeUint8, Comment: "The Task Process Result | 任务执行结果"},
		{Name: "task_task_logs", Type: field.TypeUint64, Nullable: true},
	}
	// SysTaskLogsTable holds the schema information for the "sys_task_logs" table.
	SysTaskLogsTable = &schema.Table{
		Name:       "sys_task_logs",
		Columns:    SysTaskLogsColumns,
		PrimaryKey: []*schema.Column{SysTaskLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_task_logs_sys_tasks_task_logs",
				Columns:    []*schema.Column{SysTaskLogsColumns[4]},
				RefColumns: []*schema.Column{SysTasksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SysTasksTable,
		SysTaskLogsTable,
	}
)

func init() {
	SysTasksTable.Annotation = &entsql.Annotation{
		Table: "sys_tasks",
	}
	SysTaskLogsTable.ForeignKeys[0].RefTable = SysTasksTable
	SysTaskLogsTable.Annotation = &entsql.Annotation{
		Table: "sys_task_logs",
	}
}
