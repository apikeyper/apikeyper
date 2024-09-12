'use client'

import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function OnboardingPage() {
	const router = useRouter();

	useEffect(() => {

		fetch("/api/setupWorkspace", {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
			},
		}).then((response) => {
			if (!response.ok) {
				throw new Error("Failed to setup workspace");
			}
			return response.json();
		}).then((data) => {
			localStorage.setItem("workspaceId", data.workspaceId);
			// redirect to dashboard
			router.push("/");
		});

	}, []);

	return (
		<>
			<h1>Onboarding</h1>
		</>
	);
}