const CreatePost = () => {
  return (
    <div className="bg-surface border border-outline rounded-lg p-4">
      <h3 className="font-bold text-on-surface mb-2">Create Post</h3>
      <div className="relative">
        <input
          className="bg-[#272729] border border-outline rounded-md pl-4 pr-10 py-2 w-full text-on-surface"
          placeholder="What's on your mind?"
          type="text"
        />
        <span className="material-icons-outlined absolute right-3 top-1/2 -translate-y-1/2 text-on-surface-variant">
          image
        </span>
      </div>
    </div>
  );
};

export default CreatePost;
