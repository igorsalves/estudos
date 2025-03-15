"use client";

import { Asset } from "@/models";
import { socket } from "@/socket-io";
import { useAssetStore } from "@/store";
import { useEffect } from "react";

export function AssetsSync(props: { assetsSymbols: string[] }) {
  const { assetsSymbols } = props;
  const changeAsset = useAssetStore((state) => state.changeAsset);

  useEffect(() => {
    socket.connect();

    socket.emit("joinAssets", { symbols: assetsSymbols });
    socket.on("assets/price-changed", (asset: Asset) => {
      console.log(asset);
      changeAsset(asset);
    });

    return () => {
      socket.emit("leaveAssets", { symbols: assetsSymbols });
      socket.off("assets/price-changed");
    };
  }, [assetsSymbols, changeAsset]);

  return null;
}