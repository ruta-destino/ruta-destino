export type Region = {
    id: number;
    nombre: string;
    numero: number;
};

export type Provincia = {
    id: number;
    nombre: string;
    id_region: number;
}
