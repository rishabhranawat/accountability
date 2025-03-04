import { CanActivateAuthGuard } from './common/guards/can-activate-auth-guard';
import { AuthFormComponent } from './auth-form/auth-form.component';
import { AppComponent } from './app.component';
import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DashboardComponent } from './dashboard/dashboard.component';
import { ProfileComponent } from './profile/profile.component';
import { PersonalComponent } from './personal/personal.component';

// TODO change to lazy loading where necessary
const routes: Routes = [
  { path: 'login', component: AuthFormComponent },
  { path: 'dashboard', component: DashboardComponent, canActivate: [ CanActivateAuthGuard ] },
  { path: 'personal', component: PersonalComponent },
  { path: 'profile', component: ProfileComponent },
  { path: 'logout', component: AuthFormComponent },
  { path: '', redirectTo: '/login', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
