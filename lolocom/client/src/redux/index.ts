import { combineReducers } from "redux";
import { reducer as form } from "redux-form";
import { Auth } from "./reducers/auth";
import { Cart } from "./reducers/cart";

export const reducers = combineReducers({ Auth, Cart, form });
