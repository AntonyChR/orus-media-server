import spinner from '/spinner.svg';

const Loading = () => {
    return (
        <img className='inline mb-2 animate-spin' width={20} src={spinner} />
    );
};

export default Loading;
