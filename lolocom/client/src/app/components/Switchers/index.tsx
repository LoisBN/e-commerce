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

export const AuthModeSwitcher = ({
	setter,
	current,
	texts
}: {
	setter: any;
	current: boolean;
	texts: string[];
}) => {
	return (
		<div>
			<button
				onClick={setter}
				className={"" + (current ? "--active" : "--inactive")}
			>
				{texts[0]}
			</button>
			<button
				onClick={setter}
				className={"" + (current ? "--active" : "--inactive")}
			>
				{texts[1]}
			</button>
		</div>
	);
};
