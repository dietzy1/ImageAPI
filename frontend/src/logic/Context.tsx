import { useReducer } from "react";
import { createContext, useContext } from "react";

export const defaultglobalState = {
  user: false,
};

export const GlobalStateContext = createContext(defaultglobalState);
export const DispatchStateContext = createContext<any>(undefined);

export const GlobalStateProvider = ({ children }: any) => {
  const [state, dispatch] = useReducer(
    (state: any, newValue: any) => ({ ...state, ...newValue }),
    defaultglobalState
  );
  return (
    <GlobalStateContext.Provider value={state as any}>
      <DispatchStateContext.Provider value={dispatch as any}>
        {children}
      </DispatchStateContext.Provider>
    </GlobalStateContext.Provider>
  );
};

export const useGlobalState = () => [
  useContext(GlobalStateContext),
  useContext(DispatchStateContext),
];
