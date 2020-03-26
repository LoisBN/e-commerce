import { typeofAction } from "../actions/types";

export const Auth = (state = {}, action: any) => {
	switch (action.type) {
		case typeofAction.AUTH:
			return state;
		default:
			return state;
	}
};
