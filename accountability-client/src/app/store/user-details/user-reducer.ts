import { UserState } from 'src/app/store/user-details/user-state';
import { login, setUserDetails, expireUserDetails } from './user-actions';
import { initialUserState } from './user-state';
import { Action, createReducer, on } from '@ngrx/store';

const userReducerFunc = createReducer(
    initialUserState,
    on(setUserDetails, (state, {user}) => ({...state, isProcessing: false, isAuthenticated: true, user: user})),
    on(expireUserDetails, state => ({isProcessing: false, isAuthenticated: false})),

);

export function userReducer(state: UserState | undefined, action: Action) {
  return userReducerFunc(state, action);
}
