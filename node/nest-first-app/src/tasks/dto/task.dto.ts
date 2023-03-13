import { TaskStatus } from '../task.entity';

export class CreateTaskDto {
  title: string;
  description: string;
}

export class UpdateTaskDto {
  title: string;
  description: string;
  status: TaskStatus;
}
