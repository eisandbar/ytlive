export const isEmpty = (s: string | undefined): boolean => {
  return !(s !== null && s !== undefined && s !== "");
};
