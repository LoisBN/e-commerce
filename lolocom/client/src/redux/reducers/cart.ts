import {} from "../actions/cart";
import { typeofAction } from "../actions/types";
import { Action } from "redux";

export const Cart = (state = {}, action: Action) => {
	switch (action.type) {
		case typeofAction.ADD_TO_CART:
			return {};
		case typeofAction.REMOVE_TO_CART:
			return {};
		case typeofAction.CHANGE_QUANTITY:
			return {};
		case typeofAction.CHANGE_VARIANT:
			return {};
		case typeofAction.SEND_CART:
			return {};
		default:
			return state;
	}
};
