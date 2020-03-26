import React from "react";
import {
	ImagesBanPres,
	TextedBanPres,
	FullScreenImages
} from "../components/Bans";
import { ProductPreviewPres } from "../components/Product";
import {
	ThemesPres,
	SlideCommuPres,
	SlideShowPres
} from "../components/ProductList";

export const Home = () => {
	return (
		<div>
			<ImagesBanPres images={[]} title={""} text={""} />
			<SlideShowPres images={[]} title={""} text={""} />
			<ProductPreviewPres />
			<ThemesPres themes={[]} />
			<FullScreenImages themes={[]} />
			<SlideShowPres images={[]} title={""} text={""} />
			<TextedBanPres img={""} title={""} text={""} />
			<SlideCommuPres images={[]} title={""} text={""} />
		</div>
	);
};
