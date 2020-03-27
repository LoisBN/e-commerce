import React from "react";
import c from "./index.module.css";
import { DropDownOptions, NumberChange } from "../../containers/Dropers";

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

export const ProductPreviewHomePres = ({
	marque,
	title,
	prix,
	taille,
	variants
}: {
	marque: string;
	title: string;
	prix: number[];
	taille: number[];
	variants: any;
}) => {
	const [variant, setVariant] = React.useState<number>(0);
	const [tl, setTl] = React.useState<number>(0);
	return (
		<div id={c.productpreview}>
			<h2>Zoom sur</h2>
			<div>
				<img
					alt={
						""
						//Object.keys(variants)[variant].img.title[0]
					}
					src={
						""
						//Object.keys(variants)[variant].img.title[0]
					}
				/>
				<div>
					<p>{marque}</p>
					<h3>{title}</h3>
					<div>
						<p>{prix}</p>
						<p>{prix}</p>
					</div>
					{/*
					<DropDownOptions
						list={Object.keys(taille)}
						current={Object.keys(taille)[tl]}
					/>
					<DropDownOptions
						list={Object.keys(variants)}
						current={Object.keys(variants)[variant]}
					/>
					*/}
					<NumberChange />
					<button></button>
					<button></button>
					<p>
						<a>En savoir plus</a>
					</p>
				</div>
			</div>
		</div>
	);
};
