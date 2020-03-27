import { API } from "../api";
import { typeofAction } from "./types";
import { ThunkDispatch } from "redux-thunk";

export const getProductAll = ({ ref }: { ref: string }) => async (
	dispatch: ThunkDispatch<any, undefined, any>,
	getState: any
): Promise<void> => {
	// const res = await API.;
	dispatch({
		type: typeofAction.GET_PRODUCT_ALL,
		payload: {}
	});
};

export const getProductPreview = () => async (
	dispatch: ThunkDispatch<any, undefined, any>,
	getState: any
): Promise<void> => {
	// const res = await API.;
	dispatch({
		type: typeofAction.GET_PRODUCT_PREVIEW,
		payload: {}
	});
};

export const getPromoteProducts = () => async (
	dispatch: ThunkDispatch<any, undefined, any>,
	getState: any
): Promise<void> => {
	// const res = await API.;
	dispatch({
		type: typeofAction.GET_PROMOTE_PRODUCT,
		payload: {}
	});
};

/*
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
*/
