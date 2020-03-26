import { API } from "../api";
import { typeofAction } from "./types";
import { ThunkDispatch } from "redux-thunk";

export const authentificate = () => async (
	dispatch: ThunkDispatch<any, undefined, any>,
	getState: any
): Promise<void> => {
	// const res = await API.;
	dispatch({
		type: typeofAction.AUTH,
		payload: {}
	});
};

export const signToNewsletter = () => async (
	dispatch: ThunkDispatch<any, undefined, any>,
	getState: any
): Promise<void> => {
	// const res = await API.;
	dispatch({
		type: typeofAction.NEWSLETTERS,
		payload: {}
	});
};

export const submitCart = () => async (
	dispatch: any,
	getState: any
): Promise<void> => {
	// const res = await API.;
	dispatch({
		type: typeofAction.CARTSEND,
		payload: {}
	});
};
