import { ArrowPathIcon } from '@heroicons/react/20/solid';
import { MouseEvent, ReactNode, useTransition } from 'react';

export type ButtonProps = {
  onClick: (
    event: MouseEvent<HTMLButtonElement, MouseEvent>,
  ) => void | Promise<void>;
  Icon?: ReactNode;
  children?: ReactNode;
};

const Button = (props: ButtonProps) => {
  const [isPending, startTransition] = useTransition();

  const handleButtonClick = (e: MouseEvent<HTMLButtonElement, MouseEvent>) => {
    e.preventDefault();
    startTransition(() => {
      props.onClick(e);
    });
  };

  return (
    <button
      type='button'
      disabled={isPending}
      onClick={handleButtonClick as any}
      className='inline-flex items-center gap-x-1.5 rounded-md bg-indigo-600 px-2.5 py-1.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600'
    >
      {isPending ? (
        <ArrowPathIcon className='animate-spin h-5 w-5' />
      ) : (
        props.Icon ?? null
      )}

      {props.children ? <div className='ml-3'>{props.children}</div> : null}
    </button>
  );
};

export default Button;
