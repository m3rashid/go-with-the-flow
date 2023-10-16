import { Transition } from '@headlessui/react';
import { XMarkIcon } from '@heroicons/react/20/solid';
import { Fragment, ReactNode, useState } from 'react';

import classNames from '../utils/classNames';

export type MessageProps = {
  message: string;
  description?: string;
  Icon?: ReactNode;
  hideDismiss?: boolean;
} & (
  | {
      hideUndo: false;
      undoText?: string;
      onUndo: () => void | Promise<void>;
    }
  | { hideUndo: true }
) &
  (
    | {
        hideDismiss: false;
        dismissText?: string;
        onDismiss: () => void | Promise<void>;
      }
    | { hideDismiss: true }
  );

const Message = (props: MessageProps) => {
  const [show, setShow] = useState(true);

  return (
    <div
      aria-live='assertive'
      className='pointer-events-none fixed inset-0 flex items-end px-4 py-6 sm:items-start sm:p-6'
    >
      <div className='flex w-full flex-col items-center space-y-4 sm:items-end'>
        <Transition
          show={show}
          as={Fragment}
          enter='transform ease-out duration-300 transition'
          enterFrom='translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-2'
          enterTo='translate-y-0 opacity-100 sm:translate-x-0'
          leave='transition ease-in duration-100'
          leaveFrom='opacity-100'
          leaveTo='opacity-0'
        >
          <div className='pointer-events-auto w-full max-w-sm overflow-hidden rounded-lg bg-white shadow-lg ring-1 ring-black ring-opacity-5'>
            <div className='p-4'>
              <div className='flex items-start'>
                {props.Icon ? (
                  <div className='flex-shrink-0 mr-3'>{props.Icon}</div>
                ) : null}

                <div className='w-0 flex-1 pt-0.5'>
                  <p
                    className={classNames({
                      'text-sm font-medium text-gray-900': true,
                      'w-0 flex-1': !props.hideUndo,
                    })}
                  >
                    {props.message}
                  </p>

                  {props.description ? (
                    <p className='mt-1 text-sm text-gray-500'>
                      {props.description}
                    </p>
                  ) : null}

                  <div
                    className={classNames({
                      'mt-3 flex space-x-7':
                        !props.hideUndo || !props.hideDismiss,
                    })}
                  >
                    {!props.hideUndo ? (
                      <button
                        type='button'
                        onClick={() => {
                          setShow(false);
                          props.onUndo();
                        }}
                        className='rounded-md bg-white text-sm font-medium text-indigo-600 hover:text-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2'
                      >
                        {props.undoText || 'Undo'}
                      </button>
                    ) : null}

                    {!props.hideDismiss ? (
                      <button
                        type='button'
                        onClick={() => {
                          setShow(false);
                          props.onDismiss();
                        }}
                        className='rounded-md bg-white text-sm font-medium text-gray-700 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2'
                      >
                        {props.dismissText || 'Dismiss'}
                      </button>
                    ) : null}
                  </div>
                </div>

                <div className='ml-4 flex flex-shrink-0'>
                  <button
                    type='button'
                    className='inline-flex rounded-md bg-white text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2'
                    onClick={() => setShow(false)}
                  >
                    <span className='sr-only'>Close</span>
                    <XMarkIcon className='h-5 w-5' aria-hidden='true' />
                  </button>
                </div>
              </div>
            </div>
          </div>
        </Transition>
      </div>
    </div>
  );
};

export default Message;
