import React from "react";
import { CartPres, CartPreviewPres } from "../components/Cart";

export const Cart = () => {
	return <CartPres />;
};

export const CartPreview = ({ is }: { is: boolean }) => {
	return (
		<CartPreviewPres is={is}>
			<div>
				<h3>Cart</h3>
				<i className="fas fa-times"></i>
			</div>
			<div></div>
			<div></div>
		</CartPreviewPres>
	);
};

export const CartIcon = () => {
	return <i className="fas fa-shopping-cart"></i>;
};
