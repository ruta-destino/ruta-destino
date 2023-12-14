export type Region = {
    id: number;
    nombre: string;
    numero: number;
};

export type Provincia = {
    id: number;
    nombre: string;
    id_region: number;
    nombre_region: string;
}

export type Ciudad = {
    id: number;
    nombre: string;
    id_provincia: number;
    nombre_provincia: string;
}

export type Terminal = {
    id: number;
    nombre: string;
    longitud: number;
    latitud: number;
    direccion: string;
    id_ciudad: number;
    nombre_ciudad: string;
}
