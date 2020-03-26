import React from "react";

export const ImagesBanPres = ({
	images,
	title,
	text
}: {
	images: string[];
	title?: string;
	text?: string;
}) => (
	<div>
		{images.map(img => (
			<img key={img} src={img} alt={img} />
		))}
		{title && <h2>{title}</h2>}
		{text && <p>{text}</p>}
	</div>
);

export const TextedBanPres = ({
	img,
	title,
	text
}: {
	img: string;
	title?: string;
	text?: string;
}) => (
	<div>
		<img key={img} src={img} alt={img} />
		{title && <h2>{title}</h2>}
		{text && <p>{text}</p>}
	</div>
);

export const FullScreenImages = ({
	themes
}: {
	themes: {
		title: string;
		link: { ref: string; text: string };
		img: any;
		describe: string;
	}[];
}) => (
	<div>
		{true && <button></button>}
		{themes.map(({ title, link, img, describe }) => (
			<>
				<img key={img} src={img} alt={img} />
				<h2>{title}</h2>
				<p>{describe}</p>
				<button>{link.text}</button>
			</>
		))}
	</div>
);
