import { API } from "../api";
import { typeofAction } from "./types";
import { ThunkDispatch } from "redux-thunk";

export const getCollection = ({
	which,
	filters
}: {
	which: string;
	filters?: any;
}) => async (
	dispatch: ThunkDispatch<any, undefined, any>,
	getState: any
): Promise<void> => {
	// const res = await API.;
	dispatch({
		type: typeofAction.GET_COLLECTION,
		payload: {}
	});
};

export const getCollections = ({ filter }: { filter?: string }) => async (
	dispatch: ThunkDispatch<any, undefined, any>,
	getState: any
): Promise<void> => {
	// const res = await API.;
	dispatch({
		type: typeofAction.GET_COLLECTIONS,
		payload: {}
	});
};

export const getPromoteCollections = () => async (
	dispatch: ThunkDispatch<any, undefined, any>,
	getState: any
): Promise<void> => {
	// const res = await API.;
	dispatch({
		type: typeofAction.GET_PROMOTE_COLLECTIONS,
		payload: {}
	});
};

export const getWear = () => async (
	dispatch: ThunkDispatch<any, undefined, any>,
	getState: any
): Promise<void> => {
	// const res = await API.;
	dispatch({
		type: typeofAction.GET_WEAR,
		payload: {}
	});
};
