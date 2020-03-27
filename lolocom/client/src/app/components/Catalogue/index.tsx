import React from "react";
import { ProductInCollection } from "../../containers/Product";
import c from "./index.module.css";

export const Catalogue = ({ products }: { products: any[] }) => {
	return (
		<div id={c.catalogue}>
			{products.map((productprops: any) => (
				<ProductInCollection {...productprops} />
			))}
		</div>
	);
};
