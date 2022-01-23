from typing import Any, Dict, List, Optional

import httpx

from ...client import AuthenticatedClient
from ...models.macro import Macro
from ...types import UNSET, Response


def _get_kwargs(
    *,
    client: AuthenticatedClient,
    group_id: int,
) -> Dict[str, Any]:
    url = "{}/macro/list".format(client.base_url)

    headers: Dict[str, str] = client.get_headers()
    cookies: Dict[str, Any] = client.get_cookies()

    params: Dict[str, Any] = {}
    params["groupId"] = group_id

    params = {k: v for k, v in params.items() if v is not UNSET and v is not None}

    return {
        "method": "get",
        "url": url,
        "headers": headers,
        "cookies": cookies,
        "timeout": client.get_timeout(),
        "params": params,
    }


def _parse_response(*, response: httpx.Response) -> Optional[List[Macro]]:
    if response.status_code == 200:
        response_200 = []
        _response_200 = response.json()
        for response_200_item_data in _response_200:
            response_200_item = Macro.from_dict(response_200_item_data)

            response_200.append(response_200_item)

        return response_200
    return None


def _build_response(*, response: httpx.Response) -> Response[List[Macro]]:
    return Response(
        status_code=response.status_code,
        content=response.content,
        headers=response.headers,
        parsed=_parse_response(response=response),
    )


def sync_detailed(
    *,
    client: AuthenticatedClient,
    group_id: int,
) -> Response[List[Macro]]:
    """List macros of a specific group

    Args:
        group_id (int):

    Returns:
        Response[List[Macro]]
    """

    kwargs = _get_kwargs(
        client=client,
        group_id=group_id,
    )

    response = httpx.request(
        verify=client.verify_ssl,
        **kwargs,
    )

    return _build_response(response=response)


def sync(
    *,
    client: AuthenticatedClient,
    group_id: int,
) -> Optional[List[Macro]]:
    """List macros of a specific group

    Args:
        group_id (int):

    Returns:
        Response[List[Macro]]
    """

    return sync_detailed(
        client=client,
        group_id=group_id,
    ).parsed


async def asyncio_detailed(
    *,
    client: AuthenticatedClient,
    group_id: int,
) -> Response[List[Macro]]:
    """List macros of a specific group

    Args:
        group_id (int):

    Returns:
        Response[List[Macro]]
    """

    kwargs = _get_kwargs(
        client=client,
        group_id=group_id,
    )

    async with httpx.AsyncClient(verify=client.verify_ssl) as _client:
        response = await _client.request(**kwargs)

    return _build_response(response=response)


async def asyncio(
    *,
    client: AuthenticatedClient,
    group_id: int,
) -> Optional[List[Macro]]:
    """List macros of a specific group

    Args:
        group_id (int):

    Returns:
        Response[List[Macro]]
    """

    return (
        await asyncio_detailed(
            client=client,
            group_id=group_id,
        )
    ).parsed
