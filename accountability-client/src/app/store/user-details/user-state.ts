import { User } from 'src/app/models/user.model';
export interface UserState {
  isAuthenticated: boolean;
  user?: User;
  isProcessing?: boolean;
}

export const initialUserState = { isAuthenticated: false} as UserState;
