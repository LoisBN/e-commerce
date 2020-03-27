import { API } from "../api";
import { typeofAction } from "./types";
import { ThunkDispatch } from "redux-thunk";

export const signIn = ({
	email,
	password
}: {
	email: string;
	password: string;
}) => async (
	dispatch: ThunkDispatch<any, undefined, any>,
	getState: any
): Promise<void> => {
	// const res = await API.;
	dispatch({
		type: typeofAction.SIGN_IN,
		payload: {}
	});
};
export const signUp = ({
	email,
	password,
	birth
}: {
	email: string;
	password: string;
	birth: string;
}) => async (
	dispatch: ThunkDispatch<any, undefined, any>,
	getState: any
): Promise<void> => {
	// const res = await API.;
	dispatch({
		type: typeofAction.SIGN_UP,
		payload: {}
	});
};

export const signToNewsletter = ({ mail }: { mail: string }) => async (
	dispatch: ThunkDispatch<any, undefined, any>,
	getState: any
): Promise<void> => {
	// const res = await API.;
	dispatch({
		type: typeofAction.NEWSLETTERS,
		payload: {}
	});
};

export const search = ({ text }: { text: string }) => async (
	dispatch: ThunkDispatch<any, undefined, any>,
	getState: any
): Promise<void> => {
	// const res = await API.;
	dispatch({
		type: typeofAction.SEARCH,
		payload: {}
	});
};
