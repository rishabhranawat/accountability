import { UserState } from './user-state';
import { createSelector, createFeatureSelector, select } from '@ngrx/store';

export const selectUserState = createFeatureSelector<UserState>('userStateF');

export const selectUser = createSelector(
  selectUserState,
  (state: UserState) => state.user
);

export const isAuthenticated = createSelector(
  selectUserState,
  (state: UserState) => state.isAuthenticated
);

export const isProcessing = createSelector(
  selectUserState,
  (state: UserState) => state.isProcessing
);
