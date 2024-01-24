<script lang="ts">
	export let userAvatarUrl: string;
	export let username: string;
	export let userBio: string;
	export let isLoggedInUser: boolean;
</script>

{#if isLoggedInUser}
	<div class="flex w-full flex-col gap-y-4">
		<button class="relative">
			<diV
				class="absolute h-full w-full rounded-full opacity-0 duration-200 ease-out hover:bg-gray-800 hover:opacity-80 hover:duration-200 hover:ease-in"
			>
				<p class="absolute bottom-0 w-full pb-4 text-2xl font-bold">Edit</p>
			</diV>
			<div>
				<img class="mx-auto w-2/3 rounded-full" src={userAvatarUrl} alt="User Avatar" />
			</div>
		</button>
		<h2 class="text-center text-4xl font-bold">{username}</h2>
		<p class="text-center">{userBio}</p>
	</div>
{:else}
	<div class="flex flex-col gap-y-4">
		<img class="mx-auto w-2/3 rounded-full" src={userAvatarUrl} alt="User Avatar" />
		<h2 class="text-center text-4xl font-bold">{username}</h2>
		<p class="text-center">{userBio}</p>
	</div>
{/if}
