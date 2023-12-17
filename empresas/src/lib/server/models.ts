export type Empresa = {
    id: number;
    nombre: string;
};

export type Terminal = {
    id: number;
    nombre: string;
    longitud: number;
    latitud: number;
    direccion: string;
    id_ciudad: number;
    nombre_ciudad: string;
};

export type Recorrido = {
    id: number;
    id_terminal_origen: number;
    id_terminal_destino: number;
    nombre_terminal_origen: string;
    nombre_terminal_destino: string;
    hora: number;
    minuto: number;
    dias: string;
};
