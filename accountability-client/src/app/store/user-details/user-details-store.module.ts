import { UserService } from './services/user.service';
import { UserEffects } from './user-effects';
import { StoreModule } from '@ngrx/store';
import { NgModule } from "@angular/core";
import { userReducer } from './user-reducer';
import { EffectsModule } from '@ngrx/effects';

@NgModule({
  imports: [
    StoreModule.forFeature('userStateF', userReducer),
    EffectsModule.forFeature([UserEffects])
  ],
  providers: [
    UserService
  ]
})

export class UserDetailsStoreModule {}
