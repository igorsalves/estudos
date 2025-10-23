"use client";

import { Button, TableCell, TableRow } from "flowbite-react";
import Link from "next/link";
import { useShallow } from "zustand/react/shallow";
import { useAssetStore } from "../../store";
import { Asset } from "../../models";
import { AssetShow } from "../../components/AssetShow";

export function TableAssetRow(props: Readonly<{
  asset: Asset;
  walletId: string;
}>) {
  const { asset, walletId } = props;

  const assetFound = useAssetStore(
    useShallow((state) =>
      state.assets.find((a) => a.symbol === asset.symbol)
    )
  );

  const asset_ = assetFound || asset;

  return (
    <TableRow>
    <TableCell>
      <AssetShow asset={asset_} />
    </TableCell>
    <TableCell>R$ {asset_.price}</TableCell>
    <TableCell>
      <Button
        className="w-fit"
        color="light"
        as={Link}
        href={`/assets/${asset.symbol}?wallet_id=${walletId}`}
      >
        Comprar/vender
      </Button>
    </TableCell>
  </TableRow>
  );
}