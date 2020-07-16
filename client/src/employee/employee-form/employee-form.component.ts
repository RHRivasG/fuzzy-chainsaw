import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-employee-form',
  templateUrl: './employee-form.component.html',
  styleUrls: ['./employee-form.component.scss']
})
export class EmployeeFormComponent implements OnInit {
  
  employeeForm: FormGroup

  constructor(private fb: FormBuilder) {
    this.employeeForm = fb.group({
      dummyInput: Validators.required
    })
  }

  ngOnInit(): void {
  }

}
