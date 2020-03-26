import React from "react";

export const ProductInCollectionPres = ({
	images,
	title,
	describe,
	variants,
	which
}: {
	images: string[];
	title: string;
	describe: string;
	variants: any[];
	which: any;
}) => (
	<div>
		<div>
			<img src={images[0]} alt={images[0]} />
			<img src={images[1]} alt={images[1]} />
		</div>
		<div>
			<div></div>
			<div></div>
		</div>
	</div>
);

export const ProductInCartPres = () => <></>;

export const ProductPagePres = () => <></>;

export const ProductPreviewPres = () => <></>;
