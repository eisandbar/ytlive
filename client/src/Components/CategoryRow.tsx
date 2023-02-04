import React, { ReactElement, useEffect, useState } from "react";
import { Category } from "../util/category";
import { host, categoryMinWidth } from "../util/const";
import { CategoryCard } from "./CategoryCard";
import { Grid } from "./Grid";
import { Header } from "./Header";

interface RowProps {
  gaming: boolean;
  rows?: number;
}

export const CategoryRow = ({ gaming, rows }: RowProps): ReactElement => {
  if (rows === undefined) {
    rows = 1;
  }
  const [categories, setCategories] = useState<Category[]>([]);

  useEffect(() => {
    getCategories().then(
      (res) => {},
      (rej) => {}
    );
  }, []);

  const getCategories = async (): Promise<void> => {
    try {
      const response = await fetch(
        `http://${host}/categories?gaming=${gaming.toString()}`
      );
      if (response.ok) {
        response.json().then(
          (res) => {
            setCategories(res);
          },
          (rej) => {}
        );
      }
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div className="section ">
      {gaming ? (
        <Header title="Top Gaming" url={"/categories/gaming"} gaming={true} />
      ) : (
        <Header title="Top Categories" url={"/categories/all"} gaming={false} />
      )}
      <Grid
        minWidth={categoryMinWidth}
        items={categories}
        Child={CategoryCard}
        rows={rows}
      />
      <div className="solid" />
    </div>
  );
};
