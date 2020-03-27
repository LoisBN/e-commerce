import React from "react";
import c from "./index.module.css";

export const ImagesBanPres = ({
	images,
	title,
	text
}: {
	images: string[];
	title?: string;
	text?: string;
}) => (
	<div className={`${c.imgs_ban} ${c.ban}`}>
		{images.map(img => (
			<img key={img} src={img} alt={img} />
		))}
		{title && <h2>{title}</h2>}
		{text && <p>{text}</p>}
	</div>
);

export const TextedBanImg = ({
	img,
	title,
	text
}: {
	img: string;
	title?: string;
	text?: string;
}) => (
	<div className={`${c.textedimg_ban} ${c.ban}`}>
		<img key={img} src={img} alt={img} />
		<div>
			{title && <h2>{title}</h2>}
			{text && <p>{text}</p>}
		</div>
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
		<div className={c.fullscreenimg}>
			{themes.map(({ title, link, img, describe }) => (
				<>
					<img key={img} src={img} alt={img} />
					<h2>{title}</h2>
					<p>{describe}</p>
					<button>{link.text}</button>
				</>
			))}
		</div>
	</div>
);
