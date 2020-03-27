import React from "react";

export const ColorList = ({
	colors,
	chooseColor,
	currentColor
}: {
	colors: string[];
	chooseColor: any;
	currentColor?: string;
}) => {
	return (
		<div>
			{colors.map(color => (
				<ColorItem
					key={color}
					color={color}
					isCurrent={currentColor === color}
					returnColor={chooseColor}
				/>
			))}
		</div>
	);
};

const ColorItem = ({
	color,
	returnColor,
	isCurrent
}: {
	color: string;
	returnColor: any;
	isCurrent: boolean;
}) => (
	<button
		onClick={() => {
			returnColor(color);
		}}
	></button>
);
