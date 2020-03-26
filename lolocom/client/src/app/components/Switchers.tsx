import React from "react";

export const SwitcherPres = () => {
	return <></>;
};

export const DisplayModeSwitcherPres = ({
	setter,
	current
}: {
	setter: any;
	current: boolean;
}) => {
	return (
		<div>
			<button
				onClick={setter}
				className={"" + (current ? "--active" : "--inactive")}
			></button>
			<button
				onClick={setter}
				className={"" + (current ? "--active" : "--inactive")}
			></button>
		</div>
	);
};
