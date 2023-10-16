export type ClassNames =
  | Record<string, boolean | undefined | null>
  | string
  | undefined
  | null;

const classNames = (classes: ClassNames) => {
  if (!classes) return '';
  if (typeof classes === 'string') return classes;

  return Object.entries(classes).reduce(
    (acc, [classToApply, classIfApply]) =>
      classIfApply ? `${acc} ${classToApply}` : acc,
    '',
  );
};

export default classNames;
