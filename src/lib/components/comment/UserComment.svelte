<script lang="ts">
	import type { Comment } from '$lib/data/recipe';
	import { Dropdown, DropdownItem } from 'flowbite-svelte';
	import { MoreHorizontal, ThumbsUp } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import NewCommentBar from './NewCommentBar.svelte';
	import { session } from '$src/schema';

	export let comment: Comment;
	let commentTime = new Date(Date.now() - comment.data.updatedAt.getTime()).getHours();
	let contentElem: HTMLParagraphElement;

	function isTextClamped(elm: HTMLElement) {
		return elm.scrollHeight > elm.clientHeight;
	}

	onMount(() => {
		isClamped = isTextClamped(contentElem);
	});

	$: isClamped = false;
	$: showCommentBar = false;
	$: textareaValue = '';
</script>

<div class="flex flex-row gap-x-4">
	<img src={comment.user?.avatarUrl} alt="User" class="h-8 w-8 rounded-full object-cover" />
	<div class="flex grow flex-col gap-y-2">
		<div class="flex flex-row justify-between">
			<div class="flex flex-row gap-x-2">
				<a
					class="font-bold hover:cursor-pointer hover:underline"
					href={`/user/${comment.user?.username}`}>{comment.user?.username}</a
				>
				<p class="mt-auto text-xs text-gray-500">
					{commentTime} Hour{commentTime == 1 ? '' : 's'} Ago
				</p>
			</div>
			<button>
				<MoreHorizontal
					size="26"
					class="rounded-full p-1 hover:cursor-pointer hover:bg-gray-200 dark:hover:bg-gray-800"
				/>
			</button>
			<Dropdown>
				<DropdownItem
					on:click={() => {
						navigator.clipboard.writeText(`${location.origin}/comment/${comment.data.id}`);
					}}>Copy Link</DropdownItem
				>
				<DropdownItem>Report</DropdownItem>
			</Dropdown>
		</div>
		<p bind:this={contentElem} class="line-clamp-3 pr-8">{comment.data.content}</p>
		{#if isClamped}
			<button
				class="w-fit text-xs text-gray-500 hover:cursor-pointer hover:underline hover:underline-offset-auto"
				on:click={() => {
					contentElem.classList.remove('line-clamp-3');
					isClamped = false;
				}}
			>
				Show Full Comment
			</button>
		{/if}
		<button
			class="w-min text-sm text-gray-500 hover:cursor-pointer hover:underline"
			on:click={() => {
				showCommentBar = true;
			}}>Reply</button
		>
		<div class={showCommentBar ? '' : 'hidden'}>
			<NewCommentBar
				bind:value={textareaValue}
				on:submit={async () => {
					const commentResponse = await fetch('/api/v1/comment', {
						method: 'POST',
						headers: {
							'Content-Type': 'application/json'
						},
						body: JSON.stringify({
							parentId: comment.data.id,
							recipeId: comment.data.recipeId,
							content: textareaValue
						})
					});

					if (commentResponse.ok) {
						textareaValue = '';
						showCommentBar = false;
						const values = await commentResponse.json();
						comment.children.push({
							children: [],
							data: {
								userId: values.userId,
								parentId: values.parentId,
								content: values.content,
								createdAt: new Date(Date.now()),
								id: values.id,
								recipeId: values.recipeId,
								updatedAt: new Date(Date.now())
							},
							parent: comment,
							user: values.user
						});
					} else {
						console.error('Failed to submit comment');
					}
				}}
			/>
		</div>
	</div>
</div>
