import axios from "axios";

export const API = axios.create({
	baseURL: "/service-formula"
});
