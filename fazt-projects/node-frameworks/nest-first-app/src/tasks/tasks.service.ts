import { Injectable } from '@nestjs/common';
import { Task, TaskStatus } from './task.entity';
import { v4 } from 'uuid';
import { UpdateTaskDto } from './dto/task.dto';

@Injectable()
export class TasksService {
  private tasks: Task[] = [
    {
      id: '1',
      title: 'first task',
      description: 'lalala',
      status: TaskStatus.PENDING,
    },
  ];

  getAllTasks() {
    return this.tasks;
  }
  createTask(title: string, description: string) {
    const task = {
      id: v4(),
      title,
      description,
      status: TaskStatus.PENDING,
    };

    this.tasks.push(task);

    return task;
  }

  deleteTask(id: string): Task[] {
    this.tasks = this.tasks.filter((task) => task.id !== id);
    return this.tasks;
  }

  getTaskById(id: string): Task {
    return this.tasks.find((task) => task.id === id);
  }

  updateTask(id: string, updatedFields: UpdateTaskDto): Task {
    const task = this.getTaskById(id);

    const newTask = Object.assign(task, updatedFields);
    this.tasks = this.tasks.map((item) => (item.id === id ? newTask : item));
    return newTask;
  }
}
