import { createContext } from "react";
import { AlertProps } from "./Alert";

interface ContextProps {
    showAlert(props:AlertProps): void;
}

export const AlertContext = createContext({} as ContextProps);