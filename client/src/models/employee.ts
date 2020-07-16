export interface Employee {
    nombre: string,
    apellido: string,
    nss: number,
    fecha_nac: Date,
    sexo: 'F' | 'M',
    salario: number,
    nss_sup: number | null,
    num_dep: number
}
