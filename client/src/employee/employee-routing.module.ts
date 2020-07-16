import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { EmployeeListComponent } from './employee-list/employee-list.component';
import { EmployeeResolverService } from './services/employee-resolver.service';
import { EmployeeFormComponent } from './employee-form/employee-form.component';

const routes: Routes = [
  { path: '',
    component: EmployeeListComponent,
    resolve: {
      employees: EmployeeResolverService
    },
  },
  {
    path: 'form/:id',
    component: EmployeeFormComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class EmployeeRoutingModule { }
