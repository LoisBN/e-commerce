import React from "react";
import c from "./index.module.css";
export const CartPres = () => {
	return <></>;
};

export const CartPreview = ({ unset, is }: { unset: any; is: boolean }) => (
	<div id={c.cartpreview_fullscreen} style={{ width: is ? "100vw" : 0 }}>
		<div id={c.cartpreview_out} onClick={unset} />
		<div id={c.cartpreview}>
			<div id={c.cartpreview_top}>
				<h3>Cart</h3>
				<button onClick={unset}>
					<i className="fas fa-times" />
				</button>
			</div>
			<div id={c.cartpreview_landing}></div>
			<div id={c.cartpreview_bottom}></div>
		</div>
	</div>
);
