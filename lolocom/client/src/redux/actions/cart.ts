import { typeofAction } from "./types";

export const addToCart = (productref: any, product: any, number = 1) => ({
	type: typeofAction.ADD_TO_CART,
	productref: productref,
	product: product,
	number
});

export const removeToCart = (productref: any, product: any) => ({
	type: typeofAction.REMOVE_TO_CART,
	productref: productref,
	product: product
});

export const changeVariantProduct = (productref: any, product: any) => ({
	type: typeofAction.CHANGE_VARIANT,
	productref: productref,
	product: product
});

export const changeQuantityProduct = (productref: any, product: any) => ({
	type: typeofAction.CHANGE_QUANTITY,
	productref: productref,
	product: product
});

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
