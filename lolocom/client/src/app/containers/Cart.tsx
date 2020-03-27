import React from "react";
import { CartPres, CartPreview } from "../components/Cart";

export const Cart = () => {
	return <CartPres />;
};

export const CartIcon = () => {
	const [isDisplay, setIsDisplay] = React.useState<boolean>(false);
	const switchIsDisplay = () => {
		setIsDisplay(prev => !prev);
	};
	return (
		<>
			<button onClick={switchIsDisplay}>
				<i className="fas fa-shopping-cart" />
			</button>
			<CartPreview is={isDisplay} unset={switchIsDisplay} />
		</>
	);
};
