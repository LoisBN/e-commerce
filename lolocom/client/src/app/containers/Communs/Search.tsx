import React from "react";
import { SearchBar } from "../../components/Communs/Search";

export const Search = () => {
	return <></>;
};
export const SearchIcon = () => {
	const [isDisplay, setIsDisplay] = React.useState<boolean>(false);
	const [textSearch, setTextSearch] = React.useState<string>("");
	const set = () => {
		setIsDisplay(true);
	};
	const unset = () => {
		setIsDisplay(false);
	};
	const submit = () => {};
	return (
		<>
			<button onClick={set}>
				<i className="fas fa-search" />
			</button>
			<SearchBar
				doAct={submit}
				is={isDisplay}
				text={textSearch}
				setter={setTextSearch}
				ret={unset}
			/>
		</>
	);
};
