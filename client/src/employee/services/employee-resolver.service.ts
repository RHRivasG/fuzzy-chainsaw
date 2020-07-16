import { Injectable } from '@angular/core';
import { Resolve, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import { Employee } from '../../models/employee'
import { of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class EmployeeResolverService implements Resolve<Employee[]> {

  constructor() { }
  
  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Employee[] | import("rxjs").Observable<Employee[]> | Promise<Employee[]> {
    return of([
      {
        nombre: 'Jaime',
        apellido: 'Botello',
        nss: 8886655555,
        fecha_nac: new Date(1927,11,10),
        direccion: 'Boleita',
        sexo: 'M',
        salario: 55000,
        num_dep: 1,
        nss_sup: null
      }
    ]) 
  }
}
