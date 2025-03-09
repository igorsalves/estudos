import { Badge } from "flowbite-react";
import { OrderType } from "../models";

export function OrderTypeBadge(props: Readonly<{ type: OrderType }>) {
  return (
    <Badge
      color={props.type === OrderType.BUY ? "blue" : "red"}
      className="w-fit"
    >
      {props.type === OrderType.BUY ? "Compra" : "Venda"}
    </Badge>
  );
}