import { UserState } from './user-state';
import { createSelector, createFeatureSelector } from '@ngrx/store';

export const selectStarshipsState = createFeatureSelector<UserState>('userState');

export const selectUser = createSelector(
  selectStarshipsState,
  (state: UserState) => state.user
);

export const isAuthenticated = createSelector(
  selectStarshipsState,
  (state: UserState) => state.isAuthenticated
);

export const isProcessing = createSelector(
  selectStarshipsState,
  (state: UserState) => state.isProcessing
);
