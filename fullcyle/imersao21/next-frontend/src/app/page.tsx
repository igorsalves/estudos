import { AssetShow } from "@/components/AssetShow";
import { WalletList } from "@/components/WalletList";
import { getMyWallet } from "@/queries/queries";
import { Button, Table, TableBody, TableCell, TableHead, TableHeadCell, TableRow } from "flowbite-react";
import Link from "next/link";

export default async function MyWalletListPage({
  searchParams
}: Readonly<{
  searchParams: Promise<{ wallet_id: string }>;
}>) {
  const { wallet_id } = await searchParams;

  if(!wallet_id) {
    return <WalletList />;
  }

  const wallet = await getMyWallet(wallet_id);

  if(!wallet) {
    return <WalletList />;
  }
  
  return (
    <div className="flex flex-col space-y-5 flex-grow">
      <article className="format">
        <h1>Minha carteira</h1>
      </article>
      <div className="overflow-x-auto w-full">
        <Table className="w-full max-w-full table-fixed">
          <TableHead>
            <TableHeadCell>Ativo</TableHeadCell>
            <TableHeadCell>Cotação</TableHeadCell>
            <TableHeadCell>Quantidade</TableHeadCell>
            <TableHeadCell>Comprar/Vender</TableHeadCell>
          </TableHead>
          <TableBody>
            {wallet.assets.map((walletAsset, key) => (
              <TableRow key={key}>
                <TableCell>
                  <AssetShow asset={walletAsset.asset}/>
                </TableCell>
              <TableCell>R$ {walletAsset.asset.price}</TableCell>
              <TableCell>{walletAsset.shares}</TableCell>
              <TableCell>
                <Button color="light" as={Link} href={`/assets/${walletAsset.asset.symbol}?wallet_id=${wallet_id}`}>Comprar/Vender</Button>
              </TableCell>
            </TableRow>
            ))}
            
          </TableBody>
        </Table>
      </div>
    </div>
  );
}
