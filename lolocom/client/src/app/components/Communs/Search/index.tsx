import React from "react";
import c from "./index.module.css";

export const SearchBar = ({
	text,
	setter,
	ret,
	is,
	doAct
}: {
	text: string;
	setter: any;
	ret: any;
	is: boolean;
	doAct: any;
}) => {
	const reset = () => {
		setter("");
	};
	const changeText = (e: any) => {
		setter(e.target.value);
	};
	return (
		<div id={c.searchbar} style={{ height: is ? "44px" : 0 }}>
			<input value={text} onChange={changeText} />
			<button onClick={ret}>
				<i className="fas fa-chevron-up" />
			</button>
			<button onClick={reset}>
				<i
					className="fas fa-times"
					style={{ opacity: !!text ? 1 : 0.3 }}
				/>
			</button>
			<button onClick={doAct}>
				<i className="fas fa-search" />
			</button>
		</div>
	);
};
