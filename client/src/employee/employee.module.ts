import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { EmployeeRoutingModule } from './employee-routing.module';
import { EmployeeListComponent } from './employee-list/employee-list.component';

import { MatTableModule } from '@angular/material/table'; 
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatGridListModule } from '@angular/material/grid-list';
import { EmployeeResolverService } from './services/employee-resolver.service';
import { EmployeeFormComponent } from './employee-form/employee-form.component';
import { MatFormFieldModule } from '@angular/material/form-field';
import { ReactiveFormsModule } from '@angular/forms';
import { MatInputModule } from '@angular/material/input'
import { MatSelectModule } from '@angular/material/select';

@NgModule({
  declarations: [EmployeeListComponent, EmployeeFormComponent],
  imports: [
    CommonModule,
    MatTableModule,
    MatGridListModule,
    MatButtonModule,
    MatFormFieldModule,
    MatIconModule,
    MatSelectModule,
    MatInputModule,
    ReactiveFormsModule,
    EmployeeRoutingModule
  ],
  providers: [EmployeeResolverService]
})
export class EmployeeModule { }
