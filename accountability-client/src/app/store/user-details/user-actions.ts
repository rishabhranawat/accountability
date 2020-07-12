import { User } from './../../models/user.model';

import { createAction, props } from '@ngrx/store';

export const create = createAction('[User Details] create', props<{ user: User}>());
export const login = createAction('[User Details] Login', props<{ user: User}>());
export const logout = createAction('[User Details] Logout');
export const setUserDetails = createAction('[User Details] Set User Details', props<{user: User}>());
export const expireUserDetails = createAction('[User Details] expire user details');
