const ReloadDatabase = () => {
    return (
        <div className='w-full flex flex-col'>
            <p className='text-white text-center'>
                The database appears to be empty.
            </p>
            <div className="flex justify-center">
                <button className='bg-red-800'>
                    Get information about existing files
                </button>
            </div>
        </div>
    );
};

export default ReloadDatabase;
